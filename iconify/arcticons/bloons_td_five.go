package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloonsTdFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.867 17.677c1.37-6.234 5.838-8.28 12.297-8.51C38.308 8.736 37.4 15.504 38.83 21.204c3.455 6.101.664 9.195-3.741 11.51l.067 9.787"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.466 24.48c2.224 12.593-22.116 14.468-21.94.956"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.666 28.8c.843-.079 3.71.98 5.859-.991c.76 3.308-4.347 5.737-5.859.99Zm-2.892 13.7c-1.053-3.208-1.052-6-9.274-12.036c-3.71-2.861-5.822-6.069-.968-10.796c1.117-1.088 2.225-1.356 3.335-1.991"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.669 21.898C.53 6.308 16.1 1.758 19.425 10.473m-5.558 7.204s2.048-.952 3.148-1.146c2.19-.387 6.672-.005 6.672-.005l-.043 2.565l-.98 1.707l-.203.928l-1.394.032l-.536.954l-3.734-.37c-1.08 1.63-4.574 3.633-1.187 4.573m13.905-9.83l-3.259-3.702l-2.994 3.067m-5.314 19.852c-.917 2.525-2.33 4.511-4.018 6.198M9.797 28.994c-4.785 5.09-2.038 9.354 1.72 13.506"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.526 25.436c1.134-3.225 16.637-4.262 21.94-.956m-15.544-.562l.007.785m8.69-.861l.019 1m-9.657-7.028l1.664.205l1 1.072m-4.867-.237l2.641.451l-.352 2.453"/>`),
		g.Group(children),
	)
}