package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DrDriving(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.348 21.5l18.547 7.53c2.054-2.26 5.339-3.765 7.734-1.575l8.145-4.722c.547-.753 1.05-1.848 2.44-1.232l-.319-1.643l-1.437-1.095l.342-5.612l-6.091-1.232s-7.802-1.574-13.072-1.3c0 0-4.312 4.038-5.407 6.16c0 0-10.335 2.463-10.882 4.722Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.23 16.778c3.456 1.54 13.449 3.765 15.947 3.97c0 0 1.369-6.228 2.532-8.829M24.895 29.03l-2.087 8.383C19.386 36.592 4.5 28.584 4.5 28.584c0-1.71.684-4.483 1.848-7.083m35.11-2.738l-8.281 1.985m4.907-8.349l.139 7.139M26.676 35.022l-3.868 2.391m5.905-.474c2.16-.7 2.662-2.878 2.513-5.89c-.12-2.43-1.232-3.012-2.635-2.293s-2.228 3.561-1.95 6.022c.182 1.62 1.274 2.42 2.072 2.16Zm12.804-8.634c.99-.32 1.22-1.319 1.152-2.7c-.055-1.113-.564-1.38-1.207-1.05s-1.021 1.632-.894 2.76c.083.742.584 1.109.95.99ZM31.27 32.383l9.116-5.156"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43.215 21.501l.285 3.912l-.82.593m-32.112 3.023l3.742 1.643"/>`),
		g.Group(children),
	)
}