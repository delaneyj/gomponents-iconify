package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialTumblerCircular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.377 15.59v-1.234c-.399.268-.788.4-1.166.4c-.178 0-.377-.057-.6-.166c-.134-.09-.211-.189-.234-.301c-.066-.133-.1-.422-.1-.867v-1.966h1.834v-1.233h-1.834V8.256h-1.066c-.089.467-.178.8-.267 1c-.11.244-.288.467-.533.666c-.245.201-.5.345-.767.434v1.101h.833v2.7c0 .311.044.576.134.799c.066.178.199.355.4.533c.154.156.377.289.666.4c.355.09.666.133.934.133c.311 0 .6-.033.866-.1c.312-.067.612-.178.9-.332"/><path fill="currentColor" d="M12 21c-4.963 0-9-4.037-9-9s4.037-9 9-9s9 4.037 9 9s-4.037 9-9 9zm0-16c-3.859 0-7 3.141-7 7s3.141 7 7 7s7-3.141 7-7s-3.141-7-7-7z"/>`),
		g.Group(children),
	)
}