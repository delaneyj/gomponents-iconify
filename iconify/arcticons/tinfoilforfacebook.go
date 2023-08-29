package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tinfoilforfacebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M34.2 27.052s-.527 1.102-1.238 2.286a21.695 21.695 0 0 0-2.103 4.366a10.317 10.317 0 0 0-.281 4.59c.112.485 1.318 4.721 1.43 5.206c-3.375 0-6.862-.448-10.237-.448c-.239-.758-1.755-6.007-2.384-6.297a14.574 14.574 0 0 0-2.16-.224c-.468-.098-2.78-.485-1.85-2.742c-.18-.216-.788-.469-.646-.952a1.745 1.745 0 0 0 .14-.392c-.297-.204-.756-.469-.56-.812c.3-1.063.04-1.323-.197-1.483c-.365-.021-1.327-.42-1.178-1.035a3.88 3.88 0 0 1 .814-1.26a11.47 11.47 0 0 0 1.262-2.07c.275-.637-.425-.995-.253-1.792c.467-1.52 1.009-3.728 1.009-3.728"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="m13.46 15.415l4.281.721l7.853 1.178l8.255 4.627l2.709 1.547l.19 4.75l-13.063-5.779l-12.432-3.47ZM30.424 4.5l-6.74 17.96Z"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M17.74 16.136L30.425 4.5l3.425 17.44"/>`),
		g.Group(children),
	)
}