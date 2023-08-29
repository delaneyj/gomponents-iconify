package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VrMatkalla(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.453 29.824s1.878 2.779 3.71 2.779c1.907.004 2.652-1.061 3.74-2.38l5.28-8.45a3.474 3.474 0 0 1 2.634-1.526h8.073c1.565 0 1.426 2.896-.14 2.896h-6.724a2.793 2.793 0 0 0-2.443 1.591l-1.98 3.248h5.839c1.755 0 2.901 4.629 5.323 4.629h5.063l-2.894-4.699a6.053 6.053 0 0 0 4.566-5.867c0-4.627-3.273-6.633-4.965-6.633H25.368a3.716 3.716 0 0 0-2.966 1.932l-5.082 8.409l-5.046-8.412s-1.282-1.95-2.994-1.95l-4.78-.002Z"/>`),
		g.Group(children),
	)
}