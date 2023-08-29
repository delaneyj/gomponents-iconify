package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IssueTypeFeature(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5 4.5h-.5a1 1 0 0 0 0 2h2V6A1.5 1.5 0 0 0 5 4.5zm6.5 0H11A1.5 1.5 0 0 0 9.5 6v.5h2a1 1 0 1 0 0-2z"/><path fill="currentColor" fill-rule="evenodd" d="M4 0a4 4 0 0 0-4 4v8a4 4 0 0 0 4 4h8a4 4 0 0 0 4-4V4a4 4 0 0 0-4-4H4zm.5 3a2.5 2.5 0 0 0 0 5h2.076c-.108.509-.26.902-.442 1.215c-.36.616-.891 1.008-1.68 1.346a.75.75 0 0 0 .591 1.378c.962-.412 1.806-.978 2.384-1.967c.248-.425.437-.911.571-1.47c.134.559.323 1.045.571 1.47c.578.989 1.422 1.555 2.384 1.967a.75.75 0 0 0 .59-1.378c-.787-.338-1.319-.73-1.679-1.346c-.183-.313-.334-.706-.442-1.215H11.5a2.5 2.5 0 0 0 0-5H11a3.001 3.001 0 0 0-2.83 2h-.34A3.001 3.001 0 0 0 5 3h-.5z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}