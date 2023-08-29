package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PpeMaskNNinetyFiveNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsPpeMaskN95Negative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24.294 11.044a1.001 1.001 0 0 0-.588 0l-13 4A1 1 0 0 0 10 16v1H9a5 5 0 0 0-5 5v4a5 5 0 0 0 5 5h2.456a8.991 8.991 0 0 0 4.897 3.693l7.353 2.263a.999.999 0 0 0 .588 0l7.353-2.263A8.991 8.991 0 0 0 36.544 31H39a5 5 0 0 0 5-5v-4a5 5 0 0 0-5-5h-1v-1a1 1 0 0 0-.706-.956l-13-4ZM10.482 29A8.985 8.985 0 0 1 10 26.091V19H9a3 3 0 0 0-3 3v4a3 3 0 0 0 3 3h1.482ZM39 29h-1.482A8.985 8.985 0 0 0 38 26.091V19h1a3 3 0 0 1 3 3v4a3 3 0 0 1-3 3ZM25 16v16h-2V16h2Zm6 7a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPpeMaskN95Negative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}