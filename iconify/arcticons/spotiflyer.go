package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spotiflyer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.595 32.861a21.483 21.483 0 1 0-10.73 10.733m-1.042-13.792a35.913 35.913 0 0 0-14.42-14.06c-1.377-.733-1.446-2.902-.938-4.046a2.543 2.543 0 0 1 2.14-1.437a3.262 3.262 0 0 1 .508-.001a4.517 4.517 0 0 1 1.75.54a37.378 37.378 0 0 1 16.33 16.34a4.6 4.6 0 0 1 .527 1.567m-22.535-9.68a1.877 1.877 0 0 1 .997.226A32.109 32.109 0 0 1 28.75 31.817a1.848 1.848 0 0 1-.84 2.58a2.075 2.075 0 0 1-2.688-.625a28.18 28.18 0 0 0-10.995-10.995a2.075 2.075 0 0 1-.625-2.687a1.767 1.767 0 0 1 1.583-1.065Zm-2.932 7.487a1.462 1.462 0 0 1 .753.215a24.081 24.081 0 0 1 8.267 8.267a1.323 1.323 0 0 1-.45 1.906a1.666 1.666 0 0 1-2.062-.39a21.218 21.218 0 0 0-7.27-7.272a1.678 1.678 0 0 1-.391-2.062a1.23 1.23 0 0 1 1.153-.664Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.553 28.607a7.947 7.947 0 1 0 7.947 7.947v-.01a7.937 7.937 0 0 0-7.937-7.937Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.269 32.627c2.896.944 3.99 3.232 3.848 8.343"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.529 37.725l-3.426 3.255l-3.425-3.255"/>`),
		g.Group(children),
	)
}