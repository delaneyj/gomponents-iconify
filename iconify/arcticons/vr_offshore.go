package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VrOffshore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.834 4.62c-2.569 1.209-8.487 7.101-10.325 10.652S17.16 21.92 17.16 21.92a9.535 9.535 0 0 0 2.795 2.493a12.957 12.957 0 0 0 3.4 1.133s4.004-5.44 4.532-7.026s3.098-8.536 2.947-13.9Z"/><circle cx="18.973" cy="25.903" r=".75" fill="currentColor"/><circle cx="18.299" cy="29.479" r=".75" fill="currentColor"/><circle cx="21.416" cy="31.575" r=".75" fill="currentColor"/><circle cx="24.488" cy="33.59" r=".75" fill="currentColor"/><circle cx="27.459" cy="35.655" r=".75" fill="currentColor"/><circle cx="28.517" cy="39.142" r=".75" fill="currentColor"/><circle cx="29.625" cy="42.63" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}