{
  config,
  lib,
  pkgs,
  catpix,
  ...
}:
with lib; {
  options.catpix = {
    enable = lib.mkOption {
      type = lib.types.bool;
      default = false;
      description = "Enable Catpix integration for stylix Firefox extension settings.";
    };
  };

  config = mkIf config.catpix.enable {
    home.packages = [catpix];

    home.activation.catpixGenerateSettings = lib.hm.dag.entryAfter ["writeBoundary"] ''
      ${catpix}/bin/catpix > /tmp/catpix-nix.json
      echo "done !!!"
    '';

    home-manager.firefox.extensions.settings."stylix".settings = builtins.fromJSON (builtins.readFile "/tmp/catpix-nix.json");
  };
}
