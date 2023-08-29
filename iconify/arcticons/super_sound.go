package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuperSound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.094 18.67v9.97m3.544-13.484v16.76m-7.093-10.877v5.916m-3.542-4.864v3.812M4.5 21.044v5.915m28.314-9.258v11.913m3.544-14.455v16.76m-7.093-10.877v5.915M25.723 8.833v30.334m-3.503-18.12v5.915m21.28-5.209v4.464m-3.542-7.626V29.38"/>`),
		g.Group(children),
	)
}