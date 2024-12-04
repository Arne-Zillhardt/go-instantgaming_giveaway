{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = with pkgs; [
    go
    gcc
    xorg.libX11
    xorg.libX11.dev
    xorg.xorgproto
    xorg.libXtst
    xorg.libXext
    xorg.libXi
    libxkbcommon
  ];

  shellHook = ''
    export CGO_CFLAGS="-I${pkgs.xorg.libX11.dev}/include"
    export CGO_LDFLAGS="-L${pkgs.xorg.libX11}/lib"
    export C_INCLUDE_PATH="${pkgs.xorg.libX11.dev}/include"
  '';
}
