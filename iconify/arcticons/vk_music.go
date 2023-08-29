package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VkMusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.492 18.937c-1.621-1.277-3.205-2.607-4.982-3.672c-1.792-1.073-3.678-1.9-5.752-2.27c-2.244-.403-4.36-.017-6.384.97c-1.655.806-3.16 1.855-4.668 2.9c-2.713 1.882-5.489 3.646-8.681 4.632c-2.025.625-4.1.908-6.211.966c-.102.003-.205.012-.307.018"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM12.175 29.67v-4.032m5.913 4.032v-4.032M24 32.021v-8.734m5.912-4.308V36.33m5.913-13.043v8.734"/>`),
		g.Group(children),
	)
}