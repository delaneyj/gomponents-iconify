package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phoneapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m45.004 20.266l-2.28-2.017h-.005a1.068 1.068 0 0 0-1.507.106h-.005l-4.9 5.625a1.062 1.062 0 0 0 .105 1.497l2.599 2.256c1.009.875.254 2.147-1.488 4.606a33.928 33.928 0 0 1-4.524 5.187c-1.263 1.17-2.603 2.224-3.658 1.303c0 0-.65-.567-1.295-1.133l-1.299-1.132h-.004a1.062 1.062 0 0 0-1.502.106h-.01s-5.034 5.786-5.04 5.786a1.061 1.061 0 0 0 .107 1.496l1.433 1.248"/><circle cx="19.206" cy="15.365" r="1.286" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24.092" cy="15.365" r="1.286" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="28.977" cy="15.365" r="1.286" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="19.206" cy="20.535" r="1.286" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24.092" cy="20.535" r="1.286" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="28.977" cy="20.535" r="1.286" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="19.206" cy="25.705" r="1.286" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24.092" cy="25.705" r="1.286" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="28.977" cy="25.705" r="1.286" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24.092" cy="30.875" r="1.286" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.5 24.096A21.412 21.412 0 0 1 12.651 42.19c-.98-.622-8.142 3.283-8.142 3.283s3.073-7.656 2.131-8.978a21.407 21.407 0 1 1 38.86-12.4Z"/>`),
		g.Group(children),
	)
}