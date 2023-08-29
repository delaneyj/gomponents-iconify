package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AiChatOpenAssistantChatbot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="9.636" ry="20.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.818 15.655c9.805 5.66 15.597 13.986 12.936 18.595c-2.661 4.609-12.767 3.756-22.572-1.905C9.377 26.685 3.586 18.36 6.247 13.75c1.064-1.843 3.318-2.813 6.267-2.95c4.427-.208 10.42 1.457 16.304 4.855Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.818 32.345c-9.805 5.661-19.91 6.514-22.572 1.905s3.13-12.934 12.936-18.595c5.662-3.27 11.424-4.935 15.795-4.871c3.198.046 5.652 1.018 6.777 2.966c2.66 4.609-3.13 12.934-12.936 18.595ZM20.43 21.251h7.14M19.314 24h9.372m-6.745 2.749h4.118"/>`),
		g.Group(children),
	)
}