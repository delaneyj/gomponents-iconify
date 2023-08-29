package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hibi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.5 4.5h-27a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2h27a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.723 9.245v13.851l4.426-3.362l4.425 3.362V9.245h-8.851zm22.554 11.276H26.66m0 .772V10.681h8.617v10.612m-8.617-5.825h8.617m-5.426 11.729c-.495 2.473-2.856 5.808-4.085 6.91m3.282-.527c2.064 1.197 3.566 3.029 4.68 4.149"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.048 29.439h5.963c-.995 3.71-2.455 5.186-3.623 5.908"/>`),
		g.Group(children),
	)
}