package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1006 428L744 751q-18 18-43.5 18T657 751t-18-43t18-43l227-280l-227-280q-18-18-18-43.5T657 18t43.5-18T744 18l262 323q18 18 18 43.5t-18 43.5zM367 751q-18 18-43.5 18T280 751L18 428Q0 410 0 384.5T18 341L280 18q18-18 43.5-18T367 18t18 43.5t-18 43.5L140 385l227 280q18 18 18 43t-18 43z"/>`),
		g.Group(children),
	)
}