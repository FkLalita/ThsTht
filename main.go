package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	apiKey, ok := os.LookupEnv("GEMINI_API_KEY")
	if !ok {
		log.Fatalln("Environment variable GEMINI_API_KEY not set")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	model.SetTemperature(1)
	model.SetTopK(64)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "application/json"

	// model.SafetySettings = Adjust safety settings
	// See https://ai.google.dev/gemini-api/docs/safety-settings

	session := model.StartChat()
	session.History = []*genai.Content{
		{
			Role: "user",
			Parts: []genai.Part{
				genai.Text("create something like this but with more thought provoking questions for users to vote upon\n{\n    \"category\": \"Travel\",\n    \"question_text\": \"This or That: Beach vacation or mountain getaway?\"\n    \"choice1\": \"Beach vacation\",\n    \"choice2\": \"Mountain getaway\",\n    \"meaning1\": \"Relaxing trip to a seaside location with sand and water activities\",\n    \"meaning2\": \"Trip to a mountainous region for hiking, skiing, and scenic views\"\n  },"),
			},
		},
		{
			Role: "model",
			Parts: []genai.Part{
				genai.Text("## Thought-Provoking \"This or That\" Questions for Voting\n\nHere are some examples of \"This or That\" questions with more thought-provoking options and meanings:\n\n**Category: Philosophy**\n\n```json\n{\n  \"category\": \"Philosophy\",\n  \"question_text\": \"This or That: Is it better to live a life of meaning or a life of pleasure?\",\n  \"choice1\": \"Meaningful Life\",\n  \"choice2\": \"Pleasurable Life\",\n  \"meaning1\": \"A life dedicated to achieving a purpose, making a difference, and contributing to something greater than oneself.\",\n  \"meaning2\": \"A life focused on experiencing happiness, joy, and satisfaction through personal fulfillment and sensory experiences.\"\n}\n```\n\n**Category: Technology**\n\n```json\n{\n  \"category\": \"Technology\",\n  \"question_text\": \"This or That: Is the future of work fully automated or human-centered?\",\n  \"choice1\": \"Fully Automated\",\n  \"choice2\": \"Human-Centered\",\n  \"meaning1\": \"A future where technology takes over most tasks, leaving humans to focus on higher-level thinking and creativity.\",\n  \"meaning2\": \"A future where technology complements and enhances human capabilities, fostering collaboration and innovation.\"\n}\n```\n\n**Category: Society**\n\n```json\n{\n  \"category\": \"Society\",\n  \"question_text\": \"This or That: Is it more important to focus on individual rights or collective well-being?\",\n  \"choice1\": \"Individual Rights\",\n  \"choice2\": \"Collective Well-being\",\n  \"meaning1\": \"Prioritizing the freedoms and autonomy of each individual, even if it means some potential conflict with the common good.\",\n  \"meaning2\": \"Prioritizing the health, prosperity, and security of the entire society, even if it requires some limitations on individual freedoms.\"\n}\n```\n\n**Category: Environment**\n\n```json\n{\n  \"category\": \"Environment\",\n  \"question_text\": \"This or That: Is it more effective to focus on individual actions or systemic change to address climate change?\",\n  \"choice1\": \"Individual Actions\",\n  \"choice2\": \"Systemic Change\",\n  \"meaning1\": \"Encouraging personal choices like reducing consumption and using renewable energy to make a difference.\",\n  \"meaning2\": \"Advocating for large-scale policy changes and technological advancements to address the root causes of climate change.\"\n}\n```\n\n**Category: Personal Growth**\n\n```json\n{\n  \"category\": \"Personal Growth\",\n  \"question_text\": \"This or That: Is it more important to be kind or to be right?\",\n  \"choice1\": \"Be Kind\",\n  \"choice2\": \"Be Right\",\n  \"meaning1\": \"Prioritizing empathy, understanding, and compassion over proving a point or winning an argument.\",\n  \"meaning2\": \"Strive for accuracy and logical reasoning, even if it means being assertive or challenging others' perspectives.\"\n}\n```\n\nThese are just a few examples, and you can adapt them to your specific needs and interests. \n\n**Tips for creating thought-provoking \"This or That\" questions:**\n\n* **Choose topics relevant to your audience:** Engage them by asking questions they care about and will find interesting.\n* **Offer nuanced and contrasting options:** Avoid simplistic \"black and white\" choices. Instead, provide choices that reflect complex realities and different viewpoints.\n* **Provide clear explanations for each option:** Help voters understand the underlying principles and implications of each choice.\n* **Encourage thoughtful consideration:** Ask follow-up questions or provide additional information to stimulate deeper discussion and reflection. \n"),
			},
		},
	}

	resp, err := session.SendMessage(ctx, genai.Text("INSERT_INPUT_HERE"))
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		fmt.Printf("%v\n", part)
	}
}
