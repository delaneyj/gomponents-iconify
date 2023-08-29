package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phoenix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.802 43.498h3.488c1.842 0 5.786-2.748 5.786-5.493v-5.17"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.879 6.713c5.886-3.357 13.458-2.857 18.76 1.24c5.302 4.095 7.13 10.857 4.528 16.752c-2.601 5.896-9.04 9.585-15.954 9.142c-6.914-.443-12.731-4.919-14.414-11.089"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M15.738 6.779S13.133 4.974 10.63 4.86c-.298-.014-1.963-.144-2.97-.103M8.797 43.5V8.127s-.065-2.762-1.13-3.361"/><path d="M18.498 19.49v-7.149c0-2.4-1.594-4.53-2.762-5.571"/></g><ellipse cx="24.447" cy="19.104" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.959" ry="5.872"/>`),
		g.Group(children),
	)
}