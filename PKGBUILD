pkgname=archrepo
pkgver=$(git log | wc -l)
pkgrel=1
pkgdesc="A tool to serve a arch package repo over http."
arch=(x86_64 i686)
url="https://github.com/daaku/archrepo"
source=(
  archrepo.go
  archrepo.service
  archrepo.socket
)
md5sums=($(md5sum ${source[*]} | sed -e 's/ .*//' | tr '\n' ' '))
license=('apache2')
install=install

package() {
  install -d $pkgdir/usr/bin
  go build -o $pkgdir/usr/bin/archrepo

  install -d "$pkgdir/usr/lib/systemd/system"
  install -D -m 644 "$srcdir/archrepo.service" "$pkgdir/usr/lib/systemd/system/archrepo.service"
  install -D -m 644 "$srcdir/archrepo.socket" "$pkgdir/usr/lib/systemd/system/archrepo.socket"
}
