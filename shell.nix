{ pkgs ? import <nixpkgs> { } }:

with pkgs;

mkShell {
  name = "go-dev";
  buildInputs = [
    go
    zsh
    gotools
    gopls
    go-outline
    gocode
    gopkgs
    gocode-gomod
    godef
    golint
    rnix-lsp
  ];
}

