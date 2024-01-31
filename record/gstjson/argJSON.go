package gstjson

type Source struct {
	VideoSource string `json:"videoSource"`
	AudioSource string `json:"audioSource"`
}

type Config struct {
	ArraySources []Source `json:"arraySources"`
}
