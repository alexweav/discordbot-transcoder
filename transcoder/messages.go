package transcoder

import "encoding/json"

type Message interface {
	Serialize() ([]byte, error)
}

type TranscodeCommand struct {
	FileName string `json:"fileName"`
	GuildId  string `json:"guildId"`
}

func ToTranscodeCommand(bytes []byte) (TranscodeCommand, error) {
	var command TranscodeCommand
	if err := json.Unmarshal(bytes, &command); err != nil {
		return TranscodeCommand{}, err
	}
	return command, nil
}

func (command TranscodeCommand) Serialize() ([]byte, error) {
	return json.Marshal(command)
}
