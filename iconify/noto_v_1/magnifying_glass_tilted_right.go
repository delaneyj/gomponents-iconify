package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnifyingGlassTiltedRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#78a3ad" d="m35.09 101.54l18.97-18.96c18.01 13.79 43.86 12.49 60.35-3.99c17.93-17.94 17.93-47.02 0-64.96c-17.95-17.93-47.03-17.93-64.97 0c-16.47 16.48-17.77 42.33-3.99 60.34L26.47 92.93l8.62 8.61z"/><path fill="#fff" d="M56.26 71.77c-14.16-14.18-14.16-37.15 0-51.31c14.18-14.16 37.13-14.17 51.3 0c14.16 14.16 14.16 37.13 0 51.31c-14.16 14.16-37.12 14.16-51.3 0z"/><defs><path id="notoV1MagnifyingGlassTiltedRight0" d="M4.69 123.27c4.4 4.42 10.12 5.86 12.75 3.23l25.53-25.59c6.2-6.2-9.47-22.47-15.97-15.97L1.47 110.53c-2.64 2.63-1.19 8.34 3.22 12.74z"/></defs><use fill="#f79329" href="#notoV1MagnifyingGlassTiltedRight0"/><clipPath id="notoV1MagnifyingGlassTiltedRight1"><use href="#notoV1MagnifyingGlassTiltedRight0"/></clipPath><path fill="#855c52" d="M16.77 125.77c-.18-2.04-1.51-4.45-2.37-5.81c-2.57-4.07-5.8-7.44-10.62-8.72c-4.9-1.31-5.39 2.38-4.22 6.12c1.41 4.51 4.22 7.9 8.36 10.27c.97.56 2.26 1.35 3.38 1.52c2.74.39 5.76.04 5.47-3.38z" clip-path="url(#notoV1MagnifyingGlassTiltedRight1)"/>`),
		g.Group(children),
	)
}