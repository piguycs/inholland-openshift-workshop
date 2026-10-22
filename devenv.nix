{ pkgs, ... }:

{
  packages = with pkgs; [
    go
    gopls
    sqlc
    openshift
  ];
}
