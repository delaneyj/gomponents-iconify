package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSnowfallMinimalisticBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M13 18a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm3-1.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0-3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-6 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0-3a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path d="M12 19a.995.995 0 0 0 .781-.376a.997.997 0 0 0-.182-1.425a.995.995 0 0 0-1.198 0A.999.999 0 0 0 12 19Zm3-1.5a.995.995 0 0 1-.781-.376A1 1 0 1 1 15 17.5Zm-6-2a1 1 0 1 1 0 2a1 1 0 0 1 0-2Z"/><path d="M16.286 19C19.442 19 22 16.472 22 13.353c0-2.472-1.607-4.573-3.845-5.338C17.837 5.194 15.415 3 12.476 3C9.32 3 6.762 5.528 6.762 8.647c0 .69.125 1.35.354 1.962a4.351 4.351 0 0 0-.83-.08C3.919 10.53 2 12.426 2 14.765C2 17.104 3.919 19 6.286 19h10Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}