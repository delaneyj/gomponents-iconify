package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParallelSpace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.564 38.228s-1.673-1.178-1.673-2.83V13.606s-.22-1.999 1.837-3.19L22.51 4.95c1.055-.61 1.93-.59 3.04.004l9.668 5.717s1.892 1.106 1.892 2.667v11.16c0 .394.137 1.674-.899 2.275l-9.505 6.248h-2.574v-10.2l3.97-3.096l-7.444-4.129V43.5c-3.028-1.765-4.91-3.088-8.093-5.27l-.001-.002Zm15.538-18.503l8.767 4.731m-16.212-8.86V6.032"/>`),
		g.Group(children),
	)
}