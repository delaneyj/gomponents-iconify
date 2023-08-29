package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Huaweisystemmanager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 8.4a1.56 1.56 0 0 1 .64.13l13 5.33v11.23c0 5-3.48 8.66-5.55 10.42c-3.4 2.89-7 4.09-8.1 4.09c-2.73 0-13.65-5.62-13.65-14.51V13.86l13-5.33A1.56 1.56 0 0 1 24 8.4m0-3.9a5.58 5.58 0 0 0-2.13.42L7.74 10.77a2 2 0 0 0-1.29 1.85v12.47C6.45 36.26 19.07 43.5 24 43.5s17.55-7.24 17.55-18.41V12.62a2 2 0 0 0-1.28-1.89L26.14 4.88A5.59 5.59 0 0 0 24 4.5Z"/>`),
		g.Group(children),
	)
}