package brandico

import (
	g "github.com/maragudk/gomponents"
)

const IconifyVersion = ""

func IconFromName(name string) g.Node {
	switch name {
	case "amex": return Amex()
	case "bandcamp": return Bandcamp()
	case "blogger": return Blogger()
	case "bloggerRect": return BloggerRect()
	case "box": return Box()
	case "boxRect": return BoxRect()
	case "codepen": return Codepen()
	case "deviantart": return Deviantart()
	case "diigo": return Diigo()
	case "discover": return Discover()
	case "facebook": return Facebook()
	case "facebookRect": return FacebookRect()
	case "friendfeed": return Friendfeed()
	case "friendfeedRect": return FriendfeedRect()
	case "github": return Github()
	case "githubText": return GithubText()
	case "googleplusRect": return GoogleplusRect()
	case "houzz": return Houzz()
	case "icq": return Icq()
	case "instagram": return Instagram()
	case "instagramFilled": return InstagramFilled()
	case "jabber": return Jabber()
	case "lastfm": return Lastfm()
	case "lastfmRect": return LastfmRect()
	case "linkedin": return Linkedin()
	case "linkedinRect": return LinkedinRect()
	case "mastercard": return Mastercard()
	case "odnoklassniki": return Odnoklassniki()
	case "odnoklassnikiRect": return OdnoklassnikiRect()
	case "picasa": return Picasa()
	case "skype": return Skype()
	case "tudou": return Tudou()
	case "tumblr": return Tumblr()
	case "tumblrRect": return TumblrRect()
	case "twitter": return Twitter()
	case "twitterBird": return TwitterBird()
	case "vimeo": return Vimeo()
	case "vimeoRect": return VimeoRect()
	case "visa": return Visa()
	case "vkontakteRect": return VkontakteRect()
	case "winEight": return WinEight()
	case "wordpress": return Wordpress()
	case "yandex": return Yandex()
	case "yandexRect": return YandexRect()
	case "youku": return Youku()
	default:
		panic("Unknown icon name: " + name)
	}
}
