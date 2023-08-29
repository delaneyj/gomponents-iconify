package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteHeavyCheckMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#37b34a" d="M63.792 56.915a6.876 6.876 0 0 1-6.875 6.876H6.875A6.876 6.876 0 0 1 0 56.915V6.875A6.878 6.878 0 0 1 6.875 0h50.042a6.877 6.877 0 0 1 6.875 6.875v50.04z"/><path fill="#f4f4f4" d="M53.867 14.14a4.656 4.656 0 0 0-6.562.514l-20.04 23.437l-10.781-9a4.248 4.248 0 1 0-5.447 6.519l14.444 12.06a4.223 4.223 0 0 0 3.235.946A4.654 4.654 0 0 0 31.895 47l22.483-26.3a4.659 4.659 0 0 0-.515-6.562"/>`),
		g.Group(children),
	)
}