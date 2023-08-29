package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusinessTime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 3c-.52 0-1.06.18-1.44.56c-.38.38-.56.92-.56 1.44v1H3v20h11.2c-.13-.64-.2-1.31-.2-2H5v-6.56c.59.34 1.27.56 2 .56h9.01c.56-.76 1.24-1.44 2-2H7c-1.19 0-2-.81-2-2V8h22v6c0 .16-.01.31-.04.45c.72.22 1.4.52 2.04.89V6h-9V5c0-.52-.18-1.06-.56-1.44C19.06 3.18 18.52 3 18 3h-4zm0 2h4v1h-4V5zm-5 7v3h2v-3H9zm12 0v2.46a10.153 10.153 0 0 1 2-.41V12h-2zm3 4c-4.41 0-8 3.59-8 8s3.59 8 8 8s8-3.59 8-8s-3.59-8-8-8zm0 2c3.32 0 6 2.68 6 6s-2.68 6-6 6s-6-2.68-6-6s2.68-6 6-6zm-1 1v6h5v-2h-3v-4h-2z"/>`),
		g.Group(children),
	)
}