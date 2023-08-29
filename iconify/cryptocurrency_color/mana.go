package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<defs><filter id="cryptocurrencyColorMana0"><feColorMatrix in="SourceGraphic" values="0 0 0 0 1.000000 0 0 0 0 1.000000 0 0 0 0 1.000000 0 0 0 1.000000 0"/></filter></defs><g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#FF2D55" fill-rule="nonzero"/><g filter="url(#cryptocurrencyColorMana0)"><path fill="#16141A" fill-rule="nonzero" d="m12.793 11.534l-7.045 8.454A10.912 10.912 0 0 1 5 16C5 9.923 9.923 5 16 5c6.078 0 11 4.923 11 11c0 3.36-1.507 6.369-3.883 8.387H8.883A11.511 11.511 0 0 1 7.2 22.6h12.562v-4.763l3.965 4.763H24.8l-5.043-6.05l-1.392 1.672l-5.571-6.688zM19.758 9.4a2.751 2.751 0 0 0 0 5.5a2.751 2.751 0 0 0 0-5.5zm-6.963-1.991a1.376 1.376 0 1 0 0 2.751a1.376 1.376 0 0 0 0-2.751zM9.989 25.212h12.023A10.97 10.97 0 0 1 16 27a10.97 10.97 0 0 1-6.011-1.788zm7.843-6.346l-2.426 2.909H6.639a11.056 11.056 0 0 1-.891-1.787h7.046V12.82l5.038 6.045z"/></g></g>`),
		g.Group(children),
	)
}