package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VivoGamecenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m41.846 15.218l-.047-.127c-.962-2.58-4.246-3.34-6.243-1.446l-25 23.713c-2.392 2.27-6.322.379-6.042-2.907l1.64-19.233a6.087 6.087 0 0 1 6.065-5.57H24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.175 15.026h0c.978-2.56 4.244-3.302 6.232-1.416l25.036 23.748c2.393 2.27 6.323.379 6.043-2.907l-1.64-19.233a6.087 6.087 0 0 0-6.065-5.57H24"/><circle cx="13.788" cy="24.657" r="2.758" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.534 27.415v-5.516m-2.758 2.758h5.516"/>`),
		g.Group(children),
	)
}