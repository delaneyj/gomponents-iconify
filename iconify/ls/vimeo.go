package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vimeo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M130 29h390c72 0 130 58 130 130v390c0 72-58 130-130 130H130C58 679 0 621 0 549V159C0 87 58 29 130 29zm408 216c4-20 4-41-9-57c-17-21-54-22-79-18c-20 3-89 33-113 106c42-3 64 3 60 49c-2 20-12 41-23 61c-12 23-36 69-66 36c-28-30-26-86-32-124c-4-21-8-48-14-70c-6-18-20-41-37-46c-18-5-40 3-53 11c-42 24-74 59-110 88v3c7 6 9 18 20 19c25 4 48-23 65 5c10 17 13 35 19 54c9 24 16 51 23 79c11 47 26 119 67 136c21 9 53-3 69-13c43-25 77-62 105-100c66-89 103-190 108-219z"/>`),
		g.Group(children),
	)
}