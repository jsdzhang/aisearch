from typing import List, Dict

def planner_agent_prompt(
    agent_list: list[str], agent_description: list[str], task: str
) -> str:
    prompt = f"""Break down this task into subtasks and assign to the best agent:

    TASK: {task}

    AVAILABLE AGENTS:
    {chr(10).join(f"- {agent}: {desc}" for agent, desc in zip(agent_list, agent_description))}

    RULES:
    - Only use agents from the list above
    - Skip agents not needed for the task
    - Use search agents only if updated info would help
    - Call reporter agent only once for final report
    - Each subtask must be specific

    REQUIRED FORMAT (exact JSON):
    ```json
    [
        {{
            "task": "<specific subtask>",
            "agent": "<assigned agent>"
        }}
    ]
    ```

    Respond with valid JSON only.
    """

    return prompt
