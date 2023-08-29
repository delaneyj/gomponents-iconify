package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oruxmaps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.686 10.842c.246-.016.484-.008.715.03c.223.037.44.1.646.191c3.311 1.445 3.078 9.298-.52 17.538c-3.597 8.24-9.17 11.34-12.498 9.935c-3.855-1.628-3.087-6.929.51-15.17c.246-.563.503-1.112.767-1.65c-1.167-1.562-2.338-3.125-3.47-4.645a34.258 34.258 0 0 0-2.169 4.169c-4.482 10.267-3.518 19.143 2.077 21.586c6.001 2.62 15.76-2.595 19.262-12.183c3.503-9.588 5.344-21.532-4.105-23.757h0c-1.01-.159-2.214-.367-3.718-.432"/><circle cx="24.009" cy="12.328" r="3.477" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.712 4.61a7.826 7.826 0 0 0-6.42 9.015a7.825 7.825 0 0 0 1.088 2.834c.03.05-.017-.03 0 0C20.1 20.112 23.004 24 25.773 27.7c0 0 .409.484.888.403c.478-.08.686-.67.686-.67c1.335-4.686 3.115-9.277 4.294-13.37c.142-1.013.224-2.143.087-3.032a7.826 7.826 0 0 0-9.016-6.42h0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.249 8.348c-6.015 2.746-8.52 5.611-9.028 9.458c-.508 3.848 1.84 5.758 4.928 5.11c-.468-.823-1.058-2.534-.812-4.266c.247-1.732 1.875-3.9 3.878-5.102"/>`),
		g.Group(children),
	)
}