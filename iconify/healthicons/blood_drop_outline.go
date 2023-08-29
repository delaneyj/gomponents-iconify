package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloodDropOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M24 42c6.663 0 12-5.29 12-11.737c0-5.357-3.178-11.486-6.645-16.502A72.945 72.945 0 0 0 24 6.95a72.945 72.945 0 0 0-5.355 6.811C15.178 18.777 12 24.906 12 30.263C12 36.709 17.337 42 24 42ZM22.64 5.47C19.122 9.409 10 20.56 10 30.263C10 37.85 16.268 44 24 44s14-6.15 14-13.737C38 20.56 28.878 9.409 25.36 5.47C24.52 4.53 24 4 24 4s-.52.53-1.36 1.47Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}