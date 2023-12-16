{
  inputs = {
    nixpkgs.url = github:NixOS/nixpkgs;
    flake-compat.url = "https://flakehub.com/f/edolstra/flake-compat/1.tar.gz";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = { self, nixpkgs, flake-compat, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
        aoc23 = pkgs.buildGoModule {
          pname = "aoc23";
          version = "v0.0.0";
          src = builtins.path { path = ./.; name = "aoc23"; };
          vendorHash = "sha256-m/NuMtAmCJaxxzOMidHUWoWQpCxTLqyQYQkzzr3HNfE=";
        };
      in
      {
        packages = {
          inherit aoc23;
          default = aoc23;
        };
      }
    );
}
