package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18.26 10.142A15 15 0 0 1 34 12.82a3 3 0 1 1-4 4.472a9 9 0 1 0-1 14.191V27h-5a3 3 0 1 1 0-6h8a3 3 0 0 1 3 3v8.944a3 3 0 0 1-1 2.236a15 15 0 1 1-15.74-25.038Zm8.276 1.108a13 13 0 1 0 6.13 22.44a1 1 0 0 0 .334-.746V24a1 1 0 0 0-1-1h-8a1 1 0 1 0 0 2h6a1 1 0 0 1 1 1v6a1 1 0 0 1-.4.8a11 11 0 1 1 .733-16.999a1 1 0 0 0 1.334-1.49a13 13 0 0 0-6.13-3.061Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}