package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openflashlight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.877 14.143c-.78 1.057-4.78.248-8.33-2.377s-5.27-5.838-4.489-6.895s4.29.215 7.84 2.838c3.55 2.626 5.76 5.378 4.98 6.434Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.058 4.871l-5.14 7.61a44.6 44.6 0 0 0-2.854 9.374m20.813-7.712l-5.14 7.609a32.105 32.105 0 0 1-7.991 5.35"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.064 21.855L9.302 35.978c-2.61 3.868-1.38 5.734 0 6.669c1.465.993 3.824 1.804 6.845-1.82v-.001l10.599-13.723m-4.344-15.337a14.93 14.93 0 0 0 12.805 9.29"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.064 21.855c1.097 2.956 3.841 4.843 7.682 5.248m-9.834-2.133c.868 1.728 2.105 3.44 3.249 3.987m-5.466-.781c.868 1.729 1.893 3.41 3.037 3.958m-5.008-1.107c.868 1.73 1.622 3.398 2.765 3.945m-4.674-.849c.868 1.73 1.598 3.126 2.742 3.673M8.856 36.7c.972 2.953 3.078 4.565 6.633 4.861M26.925 5.994a19.342 19.342 0 0 1 11.837 8.591"/>`),
		g.Group(children),
	)
}