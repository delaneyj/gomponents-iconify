package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WritingHandDarkSkinTone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><path id="notoV1WritingHandDarkSkinTone0" d="M.03 0h127.94v128H.03z"/></defs><clipPath id="notoV1WritingHandDarkSkinTone1"><use href="#notoV1WritingHandDarkSkinTone0"/></clipPath><path fill="#35201a" d="M25.09 96.06s-2.67 14.99 10.96 10.46c22.46-7.47 23.54.68 30.43 3.49c10.29 4.19 21.86 3.73 33.82-9.19c12.45-13.47-31.38-17.21-43.33-18.45c-11.95-1.24-30.63 13.45-30.63 13.45" clip-path="url(#notoV1WritingHandDarkSkinTone1)"/><path fill="#006ca0" d="M6.89 103.33L.73 121.09v1.12s-2.96 7.4 3.3 2.17c.42-.04.85-.19.85-.19s16.8-13.56 17.36-13.74c.56-.18 8.08-12.01 8.08-12.01s-14.25-11.39-16.6-7.78c-2.35 3.59-6.83 12.67-6.83 12.67z" clip-path="url(#notoV1WritingHandDarkSkinTone1)"/><path fill="#35201a" d="M69.93 47.4s-21.87-11.16-31.6-4.14c-8.56 6.18-15.61 9.71-19.92 12.2c-4.31 2.49-10.59 2.01-11.46 13.7c-.74 10.09-.75 18.61-2.36 22.05c-4.56 9.71 7.66 16.34 12.2-1.87c4.66-15.44 18.06-2.24 28.51-4.73c10.46-2.49 29.84-21.24 32.05-26.27c3.04-6.92-7.42-10.94-7.42-10.94z" clip-path="url(#notoV1WritingHandDarkSkinTone1)"/><path fill="#006ca0" d="M16.08 85.23S65.79 8.96 70.22 3.98c3.96-4.45.97-4.64 15.39 4.48c3.62 2.29-53.87 80.06-53.87 80.06l-16.56-1.5l.9-1.79" clip-path="url(#notoV1WritingHandDarkSkinTone1)"/><path fill="#70534a" d="M54.27 59.7S45.8 72.15 35.34 77.13c-10.46 4.98-23.41 6.97-21.42 15.44c1.99 8.47 11.95 14.45 27.89 2.49c8.47-5.48 13.45-1 20.92 2.49c7.47 3.49 8.96 9.46 27.39 6.47c18.43-2.99 20.26.43 28.39 4.48c7.97-4.98 8.96-19.92 8.96-29.39s-1.15-8.25-12.45-9.46c-8.96-2.99-33.01-14.98-39.44-19.01c-2.59-1.62-8.74-6.08-16.46 2.79" clip-path="url(#notoV1WritingHandDarkSkinTone1)"/>`),
		g.Group(children),
	)
}