package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquintingFaceWithTongue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<g fill="#fbb041" transform="matrix(1 0 0 .99955 0 .229)"><circle cx="31.749" cy="31.494" r="31.723"/><circle cx="31.749" cy="32.08" r="31.723" fill="#fbbf67"/></g><path id="emojioneV1SquintingFaceWithTongue0" fill="#633d19" d="M51.681 38.825c0 10.915-8.855 19.772-19.774 19.772c-10.911 0-19.768-8.857-19.768-19.772"/><path id="emojioneV1SquintingFaceWithTongue1" fill="#fff" d="M48.29 40.812c0 4.359-7.487 7.899-16.717 7.899c-9.223 0-16.712-3.54-16.712-7.899"/><use href="#emojioneV1SquintingFaceWithTongue0"/><path fill="#f45ba1" d="M42.641 56.54c0 4.03-4.553 7.305-10.166 7.305s-10.166-3.271-10.166-7.305l2.193-8.425c0-4.03 2.36-7.304 7.973-7.304c5.613 0 7.972 3.271 7.972 7.304l2.194 8.425"/><use href="#emojioneV1SquintingFaceWithTongue1"/><path fill="#633d19" d="M29.05 28.21c-3.03-4.98-9.168-8.04-14.989-8.217c-2.01-.062-2 2.999 0 3.061c2.793.087 5.706 1.023 8.138 2.599c-.118.202-.279 2.527-.198 2.715c-2.861.34-5.737.584-8.551 1.244c-1.952.458-1.124 3.41.831 2.952c4.557-1.069 9.28-1.037 13.837-2.106c.98-.23 1.436-1.419.932-2.248m6.33 2.247c4.558 1.069 9.283 1.037 13.838 2.106c1.956.458 2.785-2.494.831-2.952c-2.814-.66-5.69-.904-8.549-1.244c.079-.188-.08-2.513-.2-2.715c2.434-1.576 5.345-2.511 8.14-2.599c2.01-.062 2.01-3.123 0-3.061c-5.822.179-11.963 3.237-14.99 8.217c-.506.83-.048 2.02.93 2.248"/>`),
		g.Group(children),
	)
}