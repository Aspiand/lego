let pkgs = import <nixpkgs> {}; in

pkgs.mkShell {
  packages = with pkgs; [ go libgcc ];

  shellHook = ''
    [ ! -f $(pwd)/tmp ]; mkdir -p tmp/
    [ ! -f $GOPATH/bin/gin ]; go install github.com/codegangsta/gin@latest
    export PATH=$PATH:$GOPATH/bin
    export CGO_ENABLED=1 # for gorm sqlite
    # gin -i run main.go
  '';
}
