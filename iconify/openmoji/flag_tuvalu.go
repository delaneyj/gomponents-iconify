package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagTuvalu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#92d3f5" d="M5 17h62v38H5z"/><path fill="#f1b31c" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" d="m58.707 25l1.328-4l1.145 3.939L58 22.565l4-.098L58.707 25zm0 16l1.328-4l1.145 3.939L58 38.565l4-.098L58.707 41zm-6.998 4l1.328-4l1.144 3.939l-3.179-2.374l4-.098L51.709 45zm-25 7l1.328-4l1.144 3.939l-3.179-2.374l4-.098L26.709 52zm20.585-7l-1.327 4l-1.145-3.939l3.18 2.374l-4 .098L47.294 45zm-10 1l-1.327 4l-1.145-3.939l3.18 2.374l-4 .098L37.294 46zm0-7l-1.327 4l-1.145-3.939l3.18 2.374l-4 .098L37.294 39zm13.597-9.793l-.85-4.128l2.961 2.839l-3.941-.466l3.415-2.086l-1.585 3.841zm-6 2l-.85-4.128l2.961 2.839l-3.941-.466l3.415-2.086l-1.585 3.841z"/><path fill="#1e50a0" d="M5 17h31v19H5z"/><path fill="#fff" d="M9.887 18H6v2.332L32.113 36H36v-2.332L9.887 18z"/><path fill="#fff" d="M36 20.332V18h-3.887L6 33.668V36h3.887L36 20.332z"/><path fill="#fff" d="M6 24h30v6H6z"/><path fill="#fff" d="M18 18h6v18h-6z"/><path fill="#d22f27" d="M20 18h2v18h-2z"/><path fill="#d22f27" d="M6 26h30v2H6zm30 7.668L29.887 30H26l10 6v-2.332zM36 18h-3.887L24 22.868V24h2l10-6zM6 20.332L12.113 24H16L6 18v2.332zM6 36h3.887L18 31.132V30h-2L6 36z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}