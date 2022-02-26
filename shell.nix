{ pkgs ? import <nixpkgs> { } }:
pkgs.mkShell {
  shellHook = ''
    export NO_COLOR=true
    export PS1="\u@\h:\w; "
  '';

  nativeBuildInputs = with pkgs.buildPackages; [
    glfw
    go
    pkg-config
    xlibsWrapper
    xorg.libXcursor
    xorg.libXi
    xorg.libXinerama
    xorg.libXrandr
    xorg.libXxf86vm
    xorg.xinput
    xorg.xinput
  ];
}
