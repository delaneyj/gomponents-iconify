package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blossom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBlossom0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M23.266 6.008c7.734.451 7.955 8.363 7.019 11.615c9.358-2.974 12.945 1.24 13.57 3.717c1.122 5.947-4.524 7.744-7.487 7.899c5.989 7.433 3.12 11.15.936 12.08c-6.738 1.858-10.606-2.943-11.698-5.576c-2.995 6.319-8.11 6.66-10.294 6.04c-6.364-1.858-3.587-8.518-1.404-11.615c-8.984-1.487-10.294-5.885-9.826-7.898c1.497-7.434 9.358-6.195 13.101-4.647c-1.497-9.664 3.432-11.77 6.083-11.615Z"/><path d="m25 26l-4 5m-2-6l6 1m-1-7l1 7m6-2l-6 2m0 0l5 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBlossom0)"/>`),
		g.Group(children),
	)
}