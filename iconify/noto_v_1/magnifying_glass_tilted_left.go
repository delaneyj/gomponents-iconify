package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnifyingGlassTiltedLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#78a3ad" d="M101.53 92.93L82.56 73.98c13.78-18.01 12.49-43.86-3.99-60.34c-17.94-17.93-47.02-17.93-64.97 0c-17.93 17.94-17.93 47.03 0 64.96c16.48 16.47 42.34 17.78 60.35 3.99l18.97 18.96l8.61-8.62z"/><path fill="#fff" d="M20.44 71.77c-14.16-14.17-14.16-37.15 0-51.31c14.17-14.17 37.12-14.16 51.3 0c14.16 14.16 14.16 37.12 0 51.31c-14.19 14.16-37.14 14.16-51.3 0z"/><defs><path id="notoV1MagnifyingGlassTiltedLeft0" d="M126.53 110.53L101 84.94c-6.5-6.5-22.18 9.77-15.97 15.97l25.53 25.59c2.64 2.63 8.36 1.19 12.75-3.23c4.41-4.4 5.85-10.11 3.22-12.74z"/></defs><use fill="#f79329" href="#notoV1MagnifyingGlassTiltedLeft0"/><clipPath id="notoV1MagnifyingGlassTiltedLeft1"><use href="#notoV1MagnifyingGlassTiltedLeft0"/></clipPath><path fill="#855c52" d="M111.23 125.77c.18-2.04 1.51-4.45 2.37-5.81c2.57-4.07 5.8-7.44 10.62-8.72c4.9-1.31 5.39 2.38 4.22 6.12c-1.41 4.51-4.22 7.9-8.36 10.27c-.97.56-2.26 1.35-3.38 1.52c-2.74.39-5.76.04-5.47-3.38z" clip-path="url(#notoV1MagnifyingGlassTiltedLeft1)"/>`),
		g.Group(children),
	)
}