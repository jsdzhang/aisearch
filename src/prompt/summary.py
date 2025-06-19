def summary_prompt(content: str, db: list[str]) -> str:
    previous_summaries_text = "None."
    if db:
        # Enumerate for clarity if there are many previous summaries
        previous_summaries_text = "\n".join([f"{i+1}. {s}" for i, s in enumerate(db)])

    return f"""
      You are a highly capable assistant for **context-aware summarization**. Your task is to process new content, compare it against prior summaries, and generate a structured summary that captures only the **novel and distinct insights**.

      **Previously Summarized Content (for context — do NOT re-summarize):**
      {previous_summaries_text}

      **New Content to Summarize:**
      ---
      {content}
      ---

      **Instructions:**

      1. **Analyze Previous Summaries (if any):**
         - Briefly review the previous summaries to understand what has already been covered.
         - Identify existing themes or details to avoid redundancy.

      2. **Understand the New Content:**
         - Read and comprehend the main ideas, arguments, supporting details, and overall purpose.
         - Internally identify what is genuinely new, and what may overlap with prior summaries.

      3. **Generate Detailed Summary (`full_summary`):**
         - Write a clear, self-contained summary (~200–300 words) of the new content.
         - Focus strictly on new and distinct information not already covered in previous summaries.
         - Briefly mention overlaps only if necessary for context.
         - The summary should convey the key ideas, purpose, and findings of this content on its own.
         - Ignore OCR errors or irrelevant noise.

      
      4. **Craft a Distilling Short Summary (`brief_summary`):**
         - The `brief_summary` should be based on your `full_summary` and capture the core new insight introduced in this content.
         - Length: 1–3 concise sentences.
         - It must act as a unique fingerprint, clearly differentiating this entry from all prior summaries.
         - If you read this `brief_summary` later, it should immediately remind you what was uniquely added.
         - This summary will be stored for future reference, so it must be both clear and specific.

      5. **Output Format:**
         - Provide your response strictly in the following JSON format:

         ```json
            {{
            "title": "A short, clear title for this content",
            "summary": "Detailed summary (~200–300 words) focusing only on new information.",
            "brief_summary": "1–3 sentences capturing the core new insight that uniquely identifies this content.",
            "keywords": ["keyword1", "keyword2", "..."],
            "url": "source_or_placeholder"
            }}
         ```
      **Example of how to write a precise `brief_summary` (do not include this in your JSON output):**
      If earlier summaries already covered "General AI advancements" and "AI in healthcare diagnostics", and the new content discusses "Using LLMs for drug discovery", then the `brief_summary` should clearly state "LLMs applied to drug discovery", not a vague phrase like "More on AI".

      Follow these principles carefully to ensure your summaries are accurate, non-redundant, and uniquely identifiable.
      """
    
