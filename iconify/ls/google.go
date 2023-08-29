package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Google(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M82 19h491c45 0 82 37 82 82v82H553V81h-51v102H399v51h103v103h51V234h102v359c0 45-37 81-82 81H377c1-6 1-12 1-18c0-68-14-102-86-155c-20-15-66-46-66-68c0-25 7-37 45-67s66-70 66-119c0-53-23-100-64-124h59l50-53H160C97 70 38 97 0 137v-36c0-45 37-82 82-82zm166 233c10 78-25 129-84 127c-60-2-117-57-127-136s30-138 89-137c60 2 112 67 122 146zM0 543V366c29 29 70 46 121 46c7 0 14 0 21-1c-7 13-12 29-12 44c0 26 15 41 33 58h-41c-46 0-88 12-122 30zm309 131H82c-39 0-72-26-80-63c27-42 84-73 150-72c20 0 40 5 57 10c48 34 88 54 97 93c2 7 3 15 3 23v9z"/>`),
		g.Group(children),
	)
}