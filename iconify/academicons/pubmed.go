package academicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pubmed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M120.193 177.738C123.188 77.35 64.278 64 64.278 64v31.262c-23.443-11.65-45.95-13.332-45.95-13.332v255.179s98.105-11.417 98.105 84.41c0 0 36.584-153.375 248.86 26.481c0-61.593.378-216.94.378-268.285c-195.715-142.584-245.478-1.977-245.478-1.977Zm187.206 173.82l-12.376-97.654h-.448l-40.728 97.654h-17.56l-38.936-97.654h-.448l-14.176 97.654h-43.875l28.801-169.619h43.427l34.435 90.65l36.466-90.65h43.875l25.682 169.62h-44.14z"/>`),
		g.Group(children),
	)
}