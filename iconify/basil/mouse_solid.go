package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.973 2.647a.195.195 0 0 0-.223.195V8.95a.3.3 0 0 0 .3.3h5.15a.29.29 0 0 0 .293-.3c-.14-3.078-2.419-5.813-5.48-6.296a7.865 7.865 0 0 0-.04-.007Zm-1.946 0a.195.195 0 0 1 .223.195V8.95a.3.3 0 0 1-.3.3H5.8a.29.29 0 0 1-.293-.3c.14-3.078 2.419-5.813 5.48-6.296l.04-.007ZM5.8 10.75a.3.3 0 0 0-.3.3v3.876a6.5 6.5 0 1 0 13 0V11.05a.3.3 0 0 0-.3-.3H5.8Z"/>`),
		g.Group(children),
	)
}