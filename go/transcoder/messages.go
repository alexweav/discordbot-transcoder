package transcoder

import "encoding/json"

// An AMQP message.
type Message interface {
	Serialize() ([]byte, error)
}

// A command to perform an audio transcode job.
type TranscodeCommand struct {
	FileName string `json:"fileName"`
	GuildId  string `json:"guildId"`
}

// Deserializes a TranscodeCommand from JSON.
func ToTranscodeCommand(bytes []byte) (TranscodeCommand, error) {
	var command TranscodeCommand
	if err := json.Unmarshal(bytes, &command); err != nil {
		return TranscodeCommand{}, err
	}
	return command, nil
}

// Serializes a TranscodeCommand to JSON.
func (command TranscodeCommand) Serialize() ([]byte, error) {
	return json.Marshal(command)
}
