package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UzFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#1eb53a" d="M0 320h640v160H0z"/><path fill="#0099b5" d="M0 0h640v160H0z"/><path fill="#ce1126" d="M0 153.6h640v172.8H0z"/><path fill="#fff" d="M0 163.2h640v153.6H0z"/><circle cx="134.4" cy="76.8" r="57.6" fill="#fff"/><circle cx="153.6" cy="76.8" r="57.6" fill="#0099b5"/><g fill="#fff" transform="matrix(1.92 0 0 1.92 261.1 122.9)"><g id="flagUz4x30"><g id="flagUz4x31"><g id="flagUz4x32"><g id="flagUz4x33"><path id="flagUz4x34" d="M0-6L-1.9-.3L1 .7"/><use width="100%" height="100%" href="#flagUz4x34" transform="scale(-1 1)"/></g><use width="100%" height="100%" href="#flagUz4x33" transform="rotate(72)"/></g><use width="100%" height="100%" href="#flagUz4x33" transform="rotate(-72)"/><use width="100%" height="100%" href="#flagUz4x32" transform="rotate(144)"/></g><use width="100%" height="100%" y="-24" href="#flagUz4x31"/><use width="100%" height="100%" y="-48" href="#flagUz4x31"/></g><use width="100%" height="100%" x="24" href="#flagUz4x30"/><use width="100%" height="100%" x="48" href="#flagUz4x30"/><use width="100%" height="100%" x="-48" href="#flagUz4x31"/><use width="100%" height="100%" x="-24" href="#flagUz4x31"/><use width="100%" height="100%" x="-24" y="-24" href="#flagUz4x31"/></g>`),
		g.Group(children),
	)
}