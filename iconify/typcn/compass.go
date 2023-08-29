package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 5c3.859.001 7 3.142 7 7.001c0 3.858-3.141 6.998-7 6.999c-3.859 0-7-3.14-7-6.999s3.141-7 7-7.001m0-2a9 9 0 0 0 0 18a9 9 0 0 0 0-18zm4.182 4.819a.498.498 0 0 0-.491-.127L9.74 9.398a.498.498 0 0 0-.342.343l-1.707 5.951a.496.496 0 0 0 .481.637l.138-.02l5.95-1.708a.498.498 0 0 0 .342-.343l1.707-5.949a.498.498 0 0 0-.127-.49zM8.9 15.101l1.383-4.817l3.434 3.435L8.9 15.101z"/>`),
		g.Group(children),
	)
}