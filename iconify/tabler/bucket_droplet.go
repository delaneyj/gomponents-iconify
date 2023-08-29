package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BucketDroplet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m5 16l1.465 1.638a2 2 0 1 1-3.015.099L5 16zm8.737-6.263c2.299-2.3 3.23-5.095 2.081-6.245c-1.15-1.15-3.945-.217-6.244 2.082c-2.3 2.299-3.231 5.095-2.082 6.244c1.15 1.15 3.946.218 6.245-2.081z"/><path d="M7.492 11.818c.362.362.768.676 1.208.934l6.895 4.047c1.078.557 2.255-.075 3.692-1.512c1.437-1.437 2.07-2.614 1.512-3.692c-.372-.718-1.72-3.017-4.047-6.895a6.015 6.015 0 0 0-.934-1.208"/></g>`),
		g.Group(children),
	)
}