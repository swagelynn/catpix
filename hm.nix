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
      mkdir -p $HOME/.cache/catpix
      ${catpix}/bin/catpix > $HOME/.cache/catpix/stylix-settings.json
      echo "done !!!"
    '';

    home-manager.firefox.extensions.settings."stylix".settings = builtins.fromJSON (builtins.readFile "$HOME/.cache/catpix/stylix-settings.json");
  };
}
