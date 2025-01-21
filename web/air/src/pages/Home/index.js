import { Typography, Table } from '@douyinfe/semi-ui';

const Home = () => {

  const { Title } = Typography;

  const columns = [
    {
      title: '模型名称',
      dataIndex: 'name',
    },
    {
      title: '价格（每百万 Token）',
      children: [
        {
          title: '输入',
          dataIndex: 'prompt_price',
        },
        {
          title: '输出',
          dataIndex: 'output_price',
        }
      ]
    },
    {
      title: '描述',
      dataIndex: 'description',
    }
  ];

  const data = [
    {
      key: '1',
      name: 'gpt-4o ',
      prompt_price: '$2.5',
      output_price: '$10',
      description: 'OpenAI 最新的模型',
    },
    {
      key: '2',
      name: 'gpt-4o-mini',
      prompt_price: '$0.15',
      output_price: '$0.6',
      description: 'OpenAI 的小模型，替代 3.5-turbo',
    },
    {
      key: '3',
      name: 'o1-preview',
      prompt_price: '$15',
      output_price: '$60',
      description: 'OpenAI 的思考型模型，在回应之前花更多时间思考，适合复杂任务',
    },
    {
      key: '4',
      name: 'o1-mini',
      prompt_price: '$3',
      output_price: '$12',
      description: 'OpenAI 的思考型模型 mini 版',
    },
    {
      key: '5',
      name: 'claude-3.5-sonnet',
      prompt_price: '$3',
      output_price: '$15',
      description: 'Anthropic 最新的模型，代码能力强',
    },
    {
      key: '6',
      name: 'deepseek-chat',
      prompt_price: '¥1',
      output_price: '¥2',
      description: '深度求索，国内最强模型，性价比高，略逊于 gpt-4o',
    },
    {
      key: '7',
      name: 'deepseek-r1',
      prompt_price: '¥4',
      output_price: '¥16',
      description: '深度求索，国内最强思考型模型，评测分数持平 o1',
    }
  ]

  return (
    <>
      <Title heading={2} style={{ margin: '8px 0' }} >使用方式</Title>
      <p style={{ fontSize: '16px' }}>网址：<a href="https://chat.eloxt.cn">https://chat.eloxt.cn</a></p>

      <Title heading={2} style={{ margin: '8px 0' }} >设置 API Key</Title>
      <ol style={{ fontSize: '16px', lineHeight: 1.5 }}>
        <li>选择任意模型</li>
        <li>右侧打开模型设置</li>
        <li>值 - 函数 - OpenAI Manifold</li>
        <li>输入自己的 key</li>
        <li>开始聊天</li>
      </ol>

      <img src="https://s2.loli.net/2025/01/20/7rvSVUxEhA9Yl2u.png" style={{ width: '75%' }} />

      <Title heading={2} style={{ margin: '8px 0' }} >模型和定价</Title>
      <Table columns={columns} dataSource={data} pagination={false} bordered={true} />
    </>
  );
};

export default Home;