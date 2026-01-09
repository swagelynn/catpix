{
  description = "mroww";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    home-manager.url = "github:nix-community/home-manager";
    home-manager.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = {
    self,
    nixpkgs,
    home-manager,
    ...
  }: let
    system = "x86_64-linux";
    pkgs = import nixpkgs {inherit system;};

    catpix = pkgs.stdenv.mkDerivation {
      pname = "catpix";
      version = "0.1.0";

      src = ./catpix;

      buildInputs = [pkgs.go];

      buildPhase = ''
        go build -o catpix ./...
      '';

      installPhase = ''
        mkdir -p $out/bin
        cp catpix $out/bin/
      '';
    };
  in {
    homeModules = {
      catpix = {
        module = import ./hm.nix;
        pkgs = pkgs;
        catpix = catpix;
      };
    };

    homeConfigurations = {
      maddie = home-manager.lib.homeManagerConfiguration {
        inherit pkgs;
        modules = [self.homeModules.catpix.module];
        extraSpecialArgs = {
          inherit catpix;
        };
      };
    };
  };
}
