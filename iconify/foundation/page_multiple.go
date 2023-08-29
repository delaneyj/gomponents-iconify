package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M86.45 23.27h-3.475v66.91c0 .835-.677 1.513-1.513 1.513H31.987v3.475c0 .836.677 1.513 1.513 1.513h52.951c.836 0 1.513-.677 1.513-1.513V24.782a1.513 1.513 0 0 0-1.514-1.512z"/><path fill="currentColor" d="M77.988 85.193V14.807c0-.836-.677-1.513-1.513-1.513H73v66.911c0 .836-.677 1.513-1.513 1.513H22.011v3.475c0 .836.677 1.513 1.513 1.513h52.951c.836 0 1.513-.677 1.513-1.513z"/><path fill="currentColor" d="M68.013 75.218V4.832c0-.836-.677-1.513-1.513-1.513H13.55c-.836 0-1.513.677-1.513 1.513v70.386c0 .836.677 1.513 1.513 1.513H66.5c.836 0 1.513-.677 1.513-1.513z"/>`),
		g.Group(children),
	)
}