package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CartLargeMinimalisticBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3.04 2.292a.75.75 0 1 0-.498 1.416l.262.091c.667.235 1.106.39 1.429.549c.303.149.437.27.525.398c.09.132.16.314.2.677c.04.38.041.875.041 1.615V9.64h15.725c.33-1.658.485-2.5.052-3.063C20.332 6 18.815 6 17.13 6H6.492a9.029 9.029 0 0 0-.044-.738c-.053-.497-.17-.95-.452-1.362c-.284-.416-.662-.682-1.102-.899c-.412-.202-.936-.386-1.553-.603l-.301-.106Z" clip-rule="evenodd"/><path d="m20.2 12.187l.5-2.424l.024-.123H5c0 2.941.063 3.912.93 4.826c.866.914 2.26.914 5.05.914h5.302c1.561 0 2.342 0 2.893-.45c.552-.45.71-1.214 1.025-2.743Z" opacity=".5"/><path d="M7.5 18a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm9 0a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Z"/></g>`),
		g.Group(children),
	)
}