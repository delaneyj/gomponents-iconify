package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaGlobe2Fill0"><g id="evaGlobe2Fill1"><path id="evaGlobe2Fill2" fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 2a8.19 8.19 0 0 1 1.79.21a2.61 2.61 0 0 1-.78 1c-.22.17-.46.31-.7.46a4.56 4.56 0 0 0-1.85 1.67a6.49 6.49 0 0 0-.62 3.3c0 1.36 0 2.16-.95 2.87c-1.37 1.07-3.46.47-4.76-.07A8.33 8.33 0 0 1 4 12a8 8 0 0 1 8-8Zm4.89 14.32a6.79 6.79 0 0 0-.63-1.14c-.11-.16-.22-.32-.32-.49c-.39-.68-.25-1 .38-2l.1-.17a4.77 4.77 0 0 0 .58-2.43a5.42 5.42 0 0 1 .09-1c.16-.73 1.71-.93 2.67-1a7.94 7.94 0 0 1-2.86 8.28Z"/></g></g>`),
		g.Group(children),
	)
}