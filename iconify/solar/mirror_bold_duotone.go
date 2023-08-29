package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MirrorBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5.25 8.75A.75.75 0 0 1 6 9.5V19c0 .084-.003.167-.008.25h12.516A3.714 3.714 0 0 1 18.5 19V9.5a.75.75 0 0 1 1.5 0V19c0 .487.158.96.45 1.35l.9 1.2a.75.75 0 1 1-1.2.9l-.9-1.2a3.754 3.754 0 0 1-.317-.5H5.567a3.748 3.748 0 0 1-.317.5l-.9 1.2a.75.75 0 1 1-1.2-.9l.9-1.2c.292-.39.45-.863.45-1.35V9.5a.75.75 0 0 1 .75-.75Z" opacity=".5"/><path d="M12.25 17c3.866 0 7-3.358 7-7.5c0-4.142-3.134-7.5-7-7.5s-7 3.358-7 7.5c0 4.142 3.134 7.5 7 7.5Zm1.577-12.407c-.427-.194-.957-.054-1.183.312c-.227.366-.064.82.363 1.014c.792.36 1.55 1.322 1.791 2.693c.073.41.518.692.996.63c.478-.062.807-.444.735-.854c-.292-1.652-1.252-3.136-2.702-3.795Z"/></g>`),
		g.Group(children),
	)
}