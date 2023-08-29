package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Train(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M17 4.5A3.5 3.5 0 0 0 13.5 1h-7A3.5 3.5 0 0 0 3 4.5v6.955A3.5 3.5 0 0 0 5.465 14.8a15.27 15.27 0 0 0 4.535.7a15.27 15.27 0 0 0 4.535-.701A3.5 3.5 0 0 0 17 11.455V4.5ZM4 11.455V4.5A2.5 2.5 0 0 1 6.5 2h7A2.5 2.5 0 0 1 16 4.5v6.955a2.5 2.5 0 0 1-1.76 2.389A14.27 14.27 0 0 1 10 14.5a14.27 14.27 0 0 1-4.24-.656A2.5 2.5 0 0 1 4 11.455Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.5 13a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0-2a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1Zm7 2a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0-2a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1Z" clip-rule="evenodd"/><path d="M15.532 18.676a.5.5 0 0 0 .936-.351l-1.5-4a.5.5 0 0 0-.936.35l1.5 4Zm-10-4.351a.5.5 0 1 1 .936.35l-1.5 4a.5.5 0 0 1-.936-.35l1.5-4Z"/><path fill-rule="evenodd" d="M15 5a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v2.565a2 2 0 0 0 1.444 1.921C7.628 9.83 8.814 10 10 10s2.372-.171 3.556-.514A2 2 0 0 0 15 7.565V5ZM6 7.565V5a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v2.565a1 1 0 0 1-.722.96A11.746 11.746 0 0 1 10 9c-1.091 0-2.184-.158-3.278-.474A1 1 0 0 1 6 7.565Z" clip-rule="evenodd"/><path d="M5.25 17.5a.5.5 0 0 1 0-1h10a.5.5 0 0 1 0 1h-10Z"/></g>`),
		g.Group(children),
	)
}